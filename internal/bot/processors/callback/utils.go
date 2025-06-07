package callback

import (
	ctx "context"
	"fmt"
	dbmodels "pnBot/internal/db/models"
	sliceutils "pnBot/internal/sliceutils"
	"strconv"
	"strings"
)

func (cp *CallbackProcessor) getCategories() ([]dbmodels.Category, error) {
	var categories []dbmodels.Category
	context := ctx.Background()

	err := cp.dependencies.DbProvider.Find(context, &categories)
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (cp *CallbackProcessor) getUserCategories(userId int64) ([]dbmodels.Category, error) {
	var categories []dbmodels.Category

	context := ctx.Background()

	user := dbmodels.User{}
	err := cp.dependencies.DbProvider.First(context, &user, dbmodels.User{TgId: userId})
	if err != nil {
		return nil, fmt.Errorf("пользователь не найден: %w", err)
	}

	associationName := "PreferredCategories"
	out := &categories

	err = cp.dependencies.DbProvider.GetAssociation(context, user, associationName, out)
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (cp *CallbackProcessor) parseFilterData(data string) (int, []int, error) {
	parsedData := strings.Split(data, "|")
	if len(parsedData) < 3 {
		return 0, nil, fmt.Errorf("невалидный формат данных фильтра")
	}

	currentToggledCategory := parsedData[2]
	selectedCategories := parsedData[1]

	var currentToggledCategoryId int
	var err error
	if parsedData[2] == "" {
		currentToggledCategoryId = 0
	} else {
		currentToggledCategoryId, err = strconv.Atoi(currentToggledCategory)
		if err != nil {
			return 0, nil, err
		}
	}

	if selectedCategories == "" {
		return currentToggledCategoryId, []int{}, nil
	}

	parsedSelectedCategoriesIds := strings.Split(selectedCategories, ",")
	convertedSelectedCategoriesIds, err := sliceutils.StringsToInts(parsedSelectedCategoriesIds)
	if err != nil {
		return 0, nil, err
	}
	return currentToggledCategoryId, convertedSelectedCategoriesIds, nil

}

func (cp *CallbackProcessor) setSubscribed(userId int64) (bool, error) {
	context := ctx.Background()
	user := dbmodels.User{}

	where := dbmodels.User{
		TgId: userId,
	}

	if err := cp.dependencies.DbProvider.Find(context, &user, where); err != nil {
		return true, err
	}
	if user.IsSubscribed {
		return true, nil
	}

	return false, cp.dependencies.DbProvider.Update(context, where, "is_subscribed", true)
}

func (cp *CallbackProcessor) setUnsubscribed(userId int64) error {
	context := ctx.Background()

	where := dbmodels.User{
		TgId: userId,
	}

	return cp.dependencies.DbProvider.Update(context, where, "is_subscribed", false)
}

func (cp *CallbackProcessor) addPreferredCategories(userId int64, categoryIds []int) error {
	context := ctx.Background()

	user := dbmodels.User{}
	if err := cp.dependencies.DbProvider.First(context, &user, dbmodels.User{TgId: userId}); err != nil {
		return fmt.Errorf("не удалось найти пользователя: %w", err)
	}

	categories := make([]dbmodels.Category, len(categoryIds))
	for i, id := range categoryIds {
		categories[i] = dbmodels.Category{Id: uint(id)}
	}

	err := cp.dependencies.DbProvider.ReplaceAssociation(context, &user, "PreferredCategories", categories)
	if err != nil {
		return fmt.Errorf("не удалось обновить категории пользователя: %w", err)
	}

	return nil
}

func (cp *CallbackProcessor) addUserToAllCategories(userId int64) error {
	context := ctx.Background()

	user := dbmodels.User{}
	err := cp.dependencies.DbProvider.First(context, &user, dbmodels.User{TgId: userId})
	if err != nil {
		return fmt.Errorf("пользователь не найден: %w", err)
	}

	var allCategories []dbmodels.Category
	err = cp.dependencies.DbProvider.Find(context, &allCategories)
	if err != nil {
		return fmt.Errorf("не удалось получить категории: %w", err)
	}

	err = cp.dependencies.DbProvider.AddAssociation(context, &user, "PreferredCategories", allCategories)
	if err != nil {
		return fmt.Errorf("не удалось связать пользователя со всеми категориями: %w", err)
	}

	return nil
}

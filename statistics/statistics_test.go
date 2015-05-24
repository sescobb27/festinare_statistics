package statistics

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Pallinder/go-randomdata"

	"github.com/sescobb27/festinare_statistics/database"
	"github.com/sescobb27/festinare_statistics/models"
)

func shuffleCategories() []string {
	slc := make([]string, len(models.CATEGORIES))
	copy(slc, models.CATEGORIES[:])
	N := len(slc)
	for i := 0; i < N; i++ {
		// choose index uniformly in [i, N-1]
		r := i + rand.Intn(N-i)
		slc[r], slc[i] = slc[i], slc[r]
	}
	return slc
}

func FakeCategories() []*models.Category {
	tmp := shuffleCategories()
	categories := make([]*models.Category, 2)
	categories[0] = &models.Category{Name: tmp[0]}
	categories[1] = &models.Category{Name: tmp[1]}
	return categories
}

func FakeClient() *models.Client {
	categories := FakeCategories()
	discounts := make([]*models.Discount, 5)
	for i := 0; i < len(discounts); i++ {
		discounts[i] = &models.Discount{Categories: categories}
	}
	return &models.Client{
		Discounts:  discounts,
		Categories: categories,
		// Company({NAME}{NUMBER(0 to 1000)})
		Name: "Company(" + randomdata.FullName(randomdata.RandomGender) +
			string(randomdata.Number(0, 1000)) + ")",
	}
}

func FakeUser() *models.User {
	categories := FakeCategories()
	return &models.User{
		Discounts:  database.GetAvailableDiscounts(5, categories),
		Categories: FakeCategories()[:1],
	}
}

func TestUserCategoriesPreferences(t *testing.T) {
	t.Parallel()
	categoriesPreferences := database.UserCategoryPreference()
	assert.Equal(t, 3, len(categoriesPreferences), "expect 3 categories")
	sum := 0
	for _, categoryPreference := range categoriesPreferences {
		sum += categoryPreference.Count
	}
	assert.Equal(t, 1000, sum, "expect 1000 preferences")
}

func TestMain(m *testing.M) {
	fmt.Println("Droping Database")
	err := database.DB().DropDatabase()

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	fmt.Println("Inserting Clients")
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			database.ClientCollection().Insert(FakeClient())
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Inserting Users")
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			database.UserCollection().Insert(FakeUser())
			wg.Done()
		}()
	}
	wg.Wait()

	os.Exit(m.Run())
}

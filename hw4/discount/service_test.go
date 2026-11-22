package discount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type applyTable struct {
	name    string
	price   int
	percent int
	target  int
	err     error
}

type percentTable struct {
	name    string
	percent int
	err     error
}

func TestApply(t *testing.T) {
	tests := []applyTable{
		applyTable{"Regular price", 100, 10, 90, nil},
		applyTable{"Sale 0%", 100, 0, 100, nil},
		applyTable{"Sale 100%", 100, 100, 0, nil},
		applyTable{"Negative price", -100, 10, 0, ErrNegativePrice},
		applyTable{"Percent lower then 0", 100, -10, 0, ErrInvalidPercent},
		applyTable{"Percent grater then 100", 100, 101, 0, ErrInvalidPercent},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target, err := Apply(tt.price, tt.percent)
			assert.Equal(t, tt.target, target)
			assert.ErrorIs(t, err, tt.err)
		})
	}
}

func TestValidatePercent(t *testing.T) {
	tests := []percentTable{
		percentTable{"Zero percent", 0, nil},
		percentTable{"10 percent", 10, nil},
		percentTable{"100 percent", 100, nil},
		percentTable{"Invalid percent value (-1)", -1, ErrInvalidPercent},
		percentTable{"Invalid percent value (101)", 101, ErrInvalidPercent},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePercent(tt.percent)
			assert.ErrorIs(t, err, tt.err)
		})
	}
}

// Проверить несколько товаров и убедиться, что итоговая сумма рассчитывается правильно.

// Обязательно проверить:

// пустой список;
// несколько корректных товаров;
// товар с некорректной ценой;
// товар с некорректной скидкой.
func TestFinalPrice(t *testing.T) {
	emptyItemSlice := []Item{}
	someValidItems := []Item{
		Item{100, 10},
		Item{100, 1},
		Item{100, 99},
	}
	oneItemWithIncorrectPrice := []Item{
		Item{100, 10},
		Item{-100, 100},
	}
	oneItemWithIncorrectPercent := []Item{
		Item{100, -10},
		Item{100, 20},
	}

	t.Run("Empty price list", func(t *testing.T) {
		t.Parallel()
		target, err := FinalPrice(emptyItemSlice)
		assert.Equal(t, 0, target)
		assert.NoError(t, err)
	})

	t.Run("Some valid items", func(t *testing.T) {
		t.Parallel()
		target, err := FinalPrice(someValidItems)
		assert.Equal(t, 190, target)
		assert.NoError(t, err)
	})

	t.Run("One item with incorrect price", func(t *testing.T) {
		t.Parallel()
		target, err := FinalPrice(oneItemWithIncorrectPrice)
		assert.Equal(t, 0, target)
		assert.ErrorIs(t, ErrNegativePrice, err)
	})

	t.Run("One item with incorrect percent", func(t *testing.T) {
		t.Parallel()
		target, err := FinalPrice(oneItemWithIncorrectPercent)
		assert.Equal(t, 0, target)
		assert.ErrorIs(t, ErrInvalidPercent, err)
	})

}

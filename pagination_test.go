package pagination

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewInlineKeyboardPaginator(t *testing.T) {
	type args struct {
		page int
		all  int
		data string
	}
	tests := []struct {
		name string
		args args
		want []tgbotapi.InlineKeyboardButton
	}{{
		name: "when one",
		args: args{
			page: 1,
			all:  1,
			data: "data#{page}",
		},
		want: nil,
	}, {
		name: "less 5",
		args: args{
			page: 2,
			all:  4,
			data: "data#{page}",
		},
		want: []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("1", "data#1"),
			tgbotapi.NewInlineKeyboardButtonData("·2·", "data#2"),
			tgbotapi.NewInlineKeyboardButtonData("3", "data#3"),
			tgbotapi.NewInlineKeyboardButtonData("4", "data#4"),
		},
	}, {
		name: "eq 5",
		args: args{
			page: 3,
			all:  5,
			data: "data#{page}",
		},
		want: []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("1", "data#1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "data#2"),
			tgbotapi.NewInlineKeyboardButtonData("·3·", "data#3"),
			tgbotapi.NewInlineKeyboardButtonData("4", "data#4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "data#5"),
		},
	}, {
		name: "startKeyboard",
		args: args{
			page: 2,
			all:  30,
			data: "data#{page}",
		},
		want: []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("1", "data#1"),
			tgbotapi.NewInlineKeyboardButtonData("·2·", "data#2"),
			tgbotapi.NewInlineKeyboardButtonData("3", "data#3"),
			tgbotapi.NewInlineKeyboardButtonData("4 ›", "data#4"),
			tgbotapi.NewInlineKeyboardButtonData("30 »", "data#30"),
		},
	}, {
		name: "middleKeyboard",
		args: args{
			page: 15,
			all:  30,
			data: "data#{page}",
		},
		want: []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("« 1", "data#1"),
			tgbotapi.NewInlineKeyboardButtonData("‹ 14", "data#14"),
			tgbotapi.NewInlineKeyboardButtonData("·15·", "data#15"),
			tgbotapi.NewInlineKeyboardButtonData("16 ›", "data#16"),
			tgbotapi.NewInlineKeyboardButtonData("30 »", "data#30"),
		},
	}, {
		name: "finishKeyboard",
		args: args{
			page: 28,
			all:  30,
			data: "data#{page}",
		},
		want: []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("« 1", "data#1"),
			tgbotapi.NewInlineKeyboardButtonData("‹ 27", "data#27"),
			tgbotapi.NewInlineKeyboardButtonData("·28·", "data#28"),
			tgbotapi.NewInlineKeyboardButtonData("29", "data#29"),
			tgbotapi.NewInlineKeyboardButtonData("30", "data#30"),
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInlineKeyboardPaginator(tt.args.page, tt.args.all, tt.args.data)
			assert.Equal(t, tt.want, got)
		})
	}
}

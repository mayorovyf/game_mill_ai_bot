package group_admin_handlers

import (
	"gopkg.in/telebot.v3"
)

func AdminHelpHandler(c telebot.Context) error {
	helpText := `
 Команды для управления администраторами группы:

/setadmin <user_id> <звание>
— Назначить пользователя админом с указанным званием.

/removeadmin <user_id>
— Удалить пользователя из админов.

/setadmintitle <user_id> <звание>
— Изменить звание существующего админа.

/setadminrights <user_id> <ключ>=<true|false> ...
— Установить привилегии админа. Возможные ключи:
    can_edit
    can_archive
    can_manage
    can_delete

 Пример: /setadminrights 123456789 can_edit=true can_delete=false
`
	return c.Reply(helpText)
}

# Настройка GitHub и завершение работы

Product и три вопроса уже опубликованы; локальные файлы сохраняют их исходные тексты. Не создавать дубликаты:

- [product #1](https://github.com/relationskat/helpdesk-go/issues/1)
- [Хранение #2](https://github.com/relationskat/helpdesk-go/issues/2)
- [Права доступа #3](https://github.com/relationskat/helpdesk-go/issues/3)
- [Статусы #4](https://github.com/relationskat/helpdesk-go/issues/4)

1. Settings → Collaborators: предоставить преподавателю доступ. Нужен его логин; публичный репозиторий уже доступен для чтения кода, issues и PR.
2. Issues → Labels: создать review, bug, change-request, question.
3. Назначить метку question существующим issues #2–4. В product #1 после проверки отметить выполненные пункты; согласование отмечать только после ответа преподавателя.
4. Settings → Rules → Rulesets: main, Require a pull request before merging; запрет bypass, force push и удаления. Для индивидуальной работы обязательное одобрение включать при наличии рецензента. Ограничения тарифа зафиксировать и согласовать, если защита недоступна.
5. Содержательные коммиты в lab-1, отправка ветки, PR в main. Идентификаторы UR/FR вводятся с ЛР2.
6. Получить одобрение проекта и плана преподавателем в product. После этого закрыть product, слить PR, выполнить `git tag lab-1 <итоговый-SHA-main>` и `git push origin refs/tags/lab-1`.

Не закрывать product до реального одобрения. Исправления замечаний связывать с issue в сообщениях коммитов.

Фактическое ограничение на 13.11.2026: текущие GitHub credentials имеют pull=true, push=false, admin=false. Репозиторий публичный. Для отправки ветки, PR и настройки защиты требуется подключение владельца или участника с нужными правами. Product опубликован: https://github.com/relationskat/helpdesk-go/issues/1; вопросы: #2, #3, #4. Все открыты. Создание отсутствующих меток вернуло HTTP 404, назначение question — HTTP 403. Метки и защита требуют владельца.

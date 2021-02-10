from api.telegram_api import send_message


def start_handler(chat_id, arguments):
    # TODO read from .md file.
    # myteams - List the teams you are following
    # addteam - Search and select your teams. ex: /addteam {team-name}
    # removeteam - List your teams and select to remove
    text = "*Welcome to Matchday Reminder* \n Here is a list of available commands: \n `/addTeam {teamName}` Search and select a team to be notified about _\\(team name is required\\)_"
    json = send_message(chat_id=chat_id,
                        text=text,
                        options={"parse_mode": "MarkdownV2"})
    print(json)

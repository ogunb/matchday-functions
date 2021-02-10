import json


def generate_telegram_inline_keyboard_inputs(text, type, data):
    return [{
        "text": text,
        "callback_data": json.dumps({
            "data": data,
            "type": type,
        })
    }]

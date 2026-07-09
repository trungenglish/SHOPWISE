# Conversation API

Version: 1.0

---

## Purpose

Provides conversational interactions between frontend and AI runtime.

---

## Start Conversation

POST

/conversations

Returns

conversation_id

session

created_at

---

## Send Message

POST

/conversations/{id}/messages

Request

message

ui_state

conversation_context

Returns

assistant_response

ui_surfaces

actions

events

---

## Retrieve History

GET

/conversations/{id}

---

## Delete Conversation

DELETE

/conversations/{id}

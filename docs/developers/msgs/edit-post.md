# `MsgEditPost`
This message allows you to edit the message of a previously published public post.

## Structure
```json
{
  "type": "desmos/MsgEditPost",
  "value": {
    "post_id": "<ID of the post to be edited>",
    "message": "<New post message>",
    "editor": "<Desmos address of the user editing the message>",
    "edit_date": "<RFC3339-formatted date representing the date in which the post has been edited>"
  }
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `post_id` | String | ID of the post to edit |
| `message` | String | New message that should be set as the post message |
| `editor` | String | Desmos address of the user that is editing the post. This must be the same address of the original post creator. |
| `edit_date` | String | Date in RFC3339 format (e.g. `"2006-01-02T15:04:05Z07:00"`) in which the post has been edited. This must be after the original post creation date and cannot be a future date. |

## Example
### With optional data
```json
{
  "type": "desmos/MsgEditPost",
  "value": {
    "post_id": "1",
    "message": "Desmos you rock!",
    "editor": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
    "edit_date": "2020-02-05T01:00:00"
  }
}
```

## Message action
The action associated to this message is the following: 

```
edit_post
```
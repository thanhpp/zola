# Zola ~= Zalo

- [Zola ~= Zalo](#zola--zalo)
  - [Documents](#documents)
    - [Requirements](#requirements)
    - [Database](#database)
    - [Services](#services)
    - [APIs](#apis)

## Documents

### Requirements
- [GOOGLE DRIVE](https://drive.google.com/drive/folders/1ii_FZnXnlrzpcdi5AwqHDAD82lV-0S8T?usp=sharing)
- [API TESTING](https://docs.google.com/spreadsheets/d/12-7goP0F4rkHljCae2DN6iesQMPO_0gtiVl8lKsxjPA/edit?usp=sharing)
- Timestamp: Seconds from 01/01/1970 (Unix)
### Database
![Diagram](docs/Zola.png)

### Services
1. Lạc Long Quân - Manage users, posts
   1. [Postman](docs/Zola.postman_collection.json)
2. Âu Cơ - Manage chats, messages

### APIs
- [x] signup
- [x] login
- [x] logout
- [x] add_post
- [ ] get_post
- [ ] get_list_posts
- [ ] check_new_item
- [x] edit_post (missing images order)
- [x] delete_post
- [x] report
- [ ] set_comment
- [ ] get_comment
- [x] like
- [ ] edit_comment
- [ ] del_comment
- [ ] search
- [x] set_request_friend
- [ ] get_requested_friend
- [x] set_accept_friend
- [ ] get_user_friends
- [x] [change_password](https://github.com/thanhpp/zola/issues/26)
- [x] set_block_user
- [ ] set_block_diary
- [ ] get_conversation
- [ ] delete_message
- [ ] get_list_conversation
- [ ] delete_conversation

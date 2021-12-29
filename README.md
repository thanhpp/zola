# Zola ~= Zalo

- [Zola ~= Zalo](#zola--zalo)
  - [Documents](#documents)
    - [Requirements](#requirements)
    - [Database](#database)
    - [Services](#services)
    - [APIs](#apis)
      - [User](#user)
      - [Post](#post)
      - [Chat](#chat)
      - [More APIs](#more-apis)
      - [Admin APIs](#admin-apis)

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
#### User 
- [x] [signup](https://github.com/thanhpp/zola/commit/1a1bef3d247af842f8c1a16e8a4abea2c158e953)
- [x] [login](https://github.com/thanhpp/zola/issues/1)
- [x] [logout](https://github.com/thanhpp/zola/issues/3)
  
- [x] [set_request_friend](https://github.com/thanhpp/zola/issues/21)
- [ ] get_requested_friend
- [x] [set_accept_friend](https://github.com/thanhpp/zola/issues/21)
- [ ] get_user_friends

- [x] [change_password](https://github.com/thanhpp/zola/issues/26)

- [x] [set_block_user](https://github.com/thanhpp/zola/issues/19)
- [ ] set_block_diary

- [ ] [set_user_info](https://github.com/thanhpp/zola/issues/58)

#### Post
- [x] [add_post](https://github.com/thanhpp/zola/issues/5)
- [ ] get_post
  - [x] [Get post data](https://github.com/thanhpp/zola/issues/41)
  - [ ] video thumbnails
- [ ] get_list_posts
- [ ] check_new_item
- [x] [edit_post](https://github.com/thanhpp/zola/issues/7)
  - [ ] missing images order
- [x] [delete_post](https://github.com/thanhpp/zola/issues/17)
- [x] [report](https://github.com/thanhpp/zola/issues/10)
- [ ] set_comment
  - [x] [Create comment](https://github.com/thanhpp/zola/issues/28)
  - [ ] Get comment
- [ ] get_comment
- [x] [like](https://github.com/thanhpp/zola/issues/15)
- [x] [edit_comment](https://github.com/thanhpp/zola/issues/30)
- [x] [del_comment](https://github.com/thanhpp/zola/issues/31)
- [ ] search

#### Chat
- [ ] get_conversation
- [ ] delete_message
- [ ] get_list_conversation
- [ ] delete_conversation

#### More APIs
- [ ] set_official_account
- [ ] check_verify_code
- [ ] del_saved_search
- [ ] get_list_suggested_friends
- [ ] get_verify_code
- [ ] get_saved_search

#### Admin APIs
- [ ] get_admin_permission
- [ ] get_user_list
- [ ] set_role
- [ ] get_analyst_result
- [ ] set_user_state, delete_user
- [ ] get_user_basic_info

# UnOfficial 1Cak API

Kabar gembira buat kita semua, kini 1Cak ada API nya :D

### Fitur

* Dapetin postingan per-section
* Single post
* Random post (shuffle)

## Dapetin postingan per-section
```
http://api-1cak.herokuapp.com/{section}/{page id}
```

### Parameter
| Parameter        | Deskripsi                                                                     |
|------------------|-------------------------------------------------------------------------------|
| **section** _required_ | Ada 4 section, lol (_Hot_), trend (_Trending_), recent (_Vote_), legendary (_Legend_) |
| **page id** _optional_ | Untuk melihat post selanjutnya                                                |

### Contoh request

GET ```http://api-1cak.herokuapp.com/lol```

**Response** 

```
{  
    "page":{  
        "next":"1424046302"
    },
    "posts":[  
        {  
            "id":"1009826",
            "title":"Bangga Punya Walikota Seperti Ini",
            "url":"http://1cak.com/1009826",
            "img":"http://cdn16.1cak.com/posts/ddf26c6597454290c6ef7ef0f9597fc0_t.jpg",
            "votes":"268",
            "nsfw":false
        },
        {  
            "id":"1009821",
            "title":"Ini Judul",
            "url":"http://1cak.com/1009821",
            "img":"http://cdn1.1cak.tv/posts/21aaaff60ff76b067c415d642aafb9a5_t.jpg",
            "votes":"341",
            "nsfw":false
        },
        {  
            "id":"1009813",
            "title":"Wild Life!",
            "url":"http://1cak.com/1009813",
            "img":"http://cdn1.1cak.tv/posts/635fe894b99517a35dd61447a0a6d1dd_t.jpg",
            "votes":"356",
            "nsfw":false
        }
    ]
}
```

## Single Post

```
http://api-1cak.herokuapp.com/post/{id}
```

### Parameter
| Parameter        | Deskripsi                                                                     |
|------------------|-------------------------------------------------------------------------------|
| **id** _required_ | id post, numerik |

### Contoh request

GET ```http://api-1cak.herokuapp.com/post/1009826```

**Response** 

```
{  
    "id":"1009826",
    "title":"Bangga Punya Walikota Seperti Ini",
    "url":"http://1cak.com/1009826",
    "img":"http://cdn16.1cak.com/posts/ddf26c6597454290c6ef7ef0f9597fc0_t.jpg",
    "votes":"300",
    "nsfw":false
}
```

## Random Post (Shuffle)

```
http://api-1cak.herokuapp.com/random
```
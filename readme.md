# UnOfficial 1Cak API

1Cak situsnya keren, sayang gak ada API nya :(

## Resource URL
```
http://api-1cak.herokuapp.com/{section}/{page id}
```

## Parameter
| Parameter        | Deskripsi                                                                     |
|------------------|-------------------------------------------------------------------------------|
| **section** _required_ | Ada 4 section, lol (_Hot_), trend (_Trending_), recent (_Vote_), legendary (_Legend_) |
| **page id** _optional_ | Untuk melihat post selanjutnya                                                |

## Contoh request

GET ```http://api-1cak.herokuapp.com/lol```

**Response** 

```
{
    "page":{
        "next":"1416150553"
    },
    "posts":[
        {
            "id":"808972",
            "title":"Cinta Kita Ber.... Ber.... Ber.....",
            "url":"http://1cak.com/808972",
            "img":"http://cdn16.1cak.com/posts/1043ae2f43b05bcb36b957536ac10913_t.jpg",
            "votes":"150"
        },
        {
            "id":"808963",
            "title":"Boleh Dicoba Nih",
            "url":"http://1cak.com/808963",
            "img":"http://cdn16.1cak.com/posts/ba84024624d40904397781f86c11dcb7_t.jpg",
            "votes":"120"
        },
        {
            "id":"808953",
            "title":"Yang Pake N*kia Jadul Pasti Tau.",
            "url":"http://1cak.com/808953",
            "img":"http://cdn6.1cak.com/posts/0b17a77f91dc08afba3d453840f19173_t.jpg",
            "votes":"90"
        }
    ]
}
```

GET ```http://api-1cak.herokuapp.com/lol/1416150553```

**Response**

```
{
    "page":{
        "next":"1416148090"
    },
    "posts":[
        {
            "id":"808941",
            "title":"Mungkin",
            "url":"http://1cak.com/808941",
            "img":"http://cdn14.1cak.com/posts/0d60b3df3770ca20cdc1416e113c8b9e_t.jpg",
            "votes":"89"
        },
        {
            "id":"808902",
            "title":"#RESPECT Ujian Chunin Di China Karena Menghargai Tamatnya Naruto.",
            "url":"http://1cak.com/808902",
            "img":"http://cdn16.1cak.com/posts/9e430d910a0f4d813202c8eef6822015_t.jpg",
            "votes":"107"
        },
        {
            "id":"808897",
            "title":"Gitu Terusss",
            "url":"http://1cak.com/808897",
            "img":"http://cdn16.1cak.com/posts/418c3ed0ac83ea6ca86feaa57f68aa3e_t.jpg",
            "votes":"356"
        }
    ]
}
```
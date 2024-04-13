## Примеры запросов

### Работа кеша

Если сделать два успешных вызова `/user_banner` подряд, то первый вызов будет происходить из бд, а уже следущий из кеша

![cache_req](./assets/request/cache_request.png)


Чтобы убедиться в этом посмотрим логги и видим что время взятия из бд 1.768947ms, а из кеша 67.528µs:

![cache](./assets/request/cache.png)


### POST /banner

StatusCreated: `http POST http://localhost:8080/banner 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI' is_active:=true feature_id:=12 tag_ids:=[1,5] content:='{"title":"some_title","text":"some_text","url":"some_url"}'`

![post_201](./assets/request/post_201.png)

StatusBadRequest: `http POST http://localhost:8080/banner 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI' is_active:=true feature_id:=12 content:='{"title":"some_title","text":"some_text","url":"some_url"}'`

![post_400](./assets/request/post_400.png)

StatusUnauthorized: `http POST http://localhost:8080/banner is_active:=true feature_id:=12 tag_ids:=[1,5] content:='{"title":"some_title","text":"some_text","url":"some_url"}'`

![post_401](./assets/request/post_401.png)

StatusForbidden: `http POST http://localhost:8080/banner 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciJ9.9tRuwXmCywT0xIJMAFfgSW9XjsGzaFsxIbbgV3gg9Nk' is_active:=true feature_id:=12 content:='{"title":"some_title","text":"some_text","url":"some_url"}'`

![post_403](./assets/request/post_403.png)

### GET /user_banner

StatusOK: `http GET http://localhost:8080/user_banner 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI' is_active==true feature_id==12 tag_id==1`

![get_user_banner_200](./assets/request/get_user_banner_200.png)

StatusBadRequest: `http GET http://localhost:8080/user_banner 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI' is_active==true tag_id==1`

![get_user_banner_400](./assets/request/get_user_banner_400.png)

StatusUnauthorized: `http GET http://localhost:8080/user_banner is_active==true feature_id==12 tag_id==1`

![get_user_banner_401](./assets/request/get_user_banner_401.png)

StatusForbidden: `http GET http://localhost:8080/user_banner 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciJ9.9tRuwXmCywT0xIJMAFfgSW9XjsGzaFsxIbbgV3gg9Nk' is_active==true feature_id==12 tag_id==1`

![get_user_banner_403](./assets/request/get_user_banner_403.png)

StatusNotFound: `http GET http://localhost:8080/user_banner 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI' is_active==true feature_id==12 tag_id==8`

![get_user_banner_404](./assets/request/get_user_banner_404.png)

### GET /banner

StatusOK: `http GET http://localhost:8080/banner 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI'`

![get_banner_200](./assets/request/get_banner_200.png)

StatusUnauthorized: `http GET http://localhost:8080/banner is_active==true`

![get_banner_401](./assets/request/get_banner_401.png)

### PATCH /banner/{id}

StatusOK: `http PATCH http://localhost:8080/banner/1 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI' is_active:=true`

![patch_200](./assets/request/patch_200.png)

StatusBadRequest: `http PATCH http://localhost:8080/banner/1 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI'`

![patch_400](./assets/request/patch_400.png)

StatusUnauthorized: `http PATCH http://localhost:8080/banner/1`

![patch_401](./assets/request/patch_401.png)

StatusForbidden: `http PATCH http://localhost:8080/banner/1 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciJ9.9tRuwXmCywT0xIJMAFfgSW9XjsGzaFsxIbbgV3gg9Nk' is_active:=true`

![patch_403](./assets/request/patch_403.png)

StatusNotFound: `http PATCH http://localhost:8080/banner/5 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI' is_active:=true`

![patch_404](./assets/request/patch_404.png)

### DELETE /banner/{id}

StatusNoContent: `http DELETE http://localhost:8080/banner/1 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI'`

![delete_204](./assets/request/delete_204.png)

StatusUnauthorized: `http DELETE http://localhost:8080/banner/1`

![delete_401](./assets/request/delete_401.png)

StatusForbidden: `http DELETE http://localhost:8080/banner/1 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciJ9.9tRuwXmCywT0xIJMAFfgSW9XjsGzaFsxIbbgV3gg9Nk'`

![delete_403](./assets/request/delete_403.png)

StatusNotFound: `http DELETE http://localhost:8080/banner/1 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI'`

![delete_404](./assets/request/delete_404.png)

### GET /token

StatusCreated: `http GET "http://localhost:8080/token"`

![token](./assets/request/token.png)


## END_TO_END

Тесты выполняются для все сценариев статуса ответа для каждого эндпоинта. Использован пакет `github.com/gavv/httpexpect/v2`. Рекомендуется делать в пустой базе данных, иначе можно получить неожиданные результаты.

![testing](./assets/end_to_end/test.png)


## STRESS TESTS

Время подсчитывается путем суммирования времени для каждого запроса, даже если они выполнялись одновременно. Время высчитывается между подачей запроса и получением ответа. Подается POST запрос на создание баннера, возвращается ID каждого баннера, затем по этим ID происходит удаление каждого баннера. Рекомендуется делать в пустой базе данных, иначе можно получить неожиданные результаты.

### 1000RPS

Результаты тестов для 1000 запросов в секунду, все из которых успешны

![1000RPS](./assets/stress/1000RPS.png)

Результаты тестов для 1000 запросов в секунду, все из которых успешны. На любом промежутке логов хорошее время выполнения
```
time=2024-04-13T19:36:57.353+03:00 level=INFO msg="Request status" Path=/banner/980 Method=DELETE status_code=204 elapsed_time=7.36237ms
time=2024-04-13T19:36:57.356+03:00 level=INFO msg="Request status" Path=/banner/988 Method=DELETE status_code=204 elapsed_time=6.034278ms
time=2024-04-13T19:36:57.356+03:00 level=INFO msg="Request status" Path=/banner/987 Method=DELETE status_code=204 elapsed_time=4.976927ms
time=2024-04-13T19:36:57.356+03:00 level=INFO msg="Request status" Path=/banner/985 Method=DELETE status_code=204 elapsed_time=7.501203ms
time=2024-04-13T19:36:57.360+03:00 level=INFO msg="Request status" Path=/banner/989 Method=DELETE status_code=204 elapsed_time=4.816323ms
time=2024-04-13T19:36:57.360+03:00 level=INFO msg="Request status" Path=/banner/991 Method=DELETE status_code=204 elapsed_time=5.901867ms
time=2024-04-13T19:36:57.360+03:00 level=INFO msg="Request status" Path=/banner/986 Method=DELETE status_code=204 elapsed_time=7.258384ms
time=2024-04-13T19:36:57.364+03:00 level=INFO msg="Request status" Path=/banner/994 Method=DELETE status_code=204 elapsed_time=5.861009ms
time=2024-04-13T19:36:57.364+03:00 level=INFO msg="Request status" Path=/banner/992 Method=DELETE status_code=204 elapsed_time=4.773542ms
time=2024-04-13T19:36:57.364+03:00 level=INFO msg="Request status" Path=/banner/990 Method=DELETE status_code=204 elapsed_time=7.204942ms
time=2024-04-13T19:36:57.367+03:00 level=INFO msg="Request status" Path=/banner/996 Method=DELETE status_code=204 elapsed_time=5.043453ms
time=2024-04-13T19:36:57.368+03:00 level=INFO msg="Request status" Path=/banner/993 Method=DELETE status_code=204 elapsed_time=7.371568ms
time=2024-04-13T19:36:57.368+03:00 level=INFO msg="Request status" Path=/banner/998 Method=DELETE status_code=204 elapsed_time=6.284231ms
time=2024-04-13T19:36:57.371+03:00 level=INFO msg="Request status" Path=/banner/1000 Method=DELETE status_code=204 elapsed_time=5.106411ms
```
### 2000RPS

![2000RPS](./assets/stress/2000RPS.png)

Тут невооруженным глазом видны проблемы, если посмотреть что происходит в логах, то первые операции выполняются быстро:

```
time=2024-04-13T19:47:23.320+03:00 level=INFO msg="Request status" Path=/banner/1006 Method=DELETE status_code=204 elapsed_time=7.706581ms
time=2024-04-13T19:47:23.322+03:00 level=INFO msg="Request status" Path=/banner/1067 Method=DELETE status_code=204 elapsed_time=7.122025ms
time=2024-04-13T19:47:23.323+03:00 level=INFO msg="Request status" Path=/banner/1010 Method=DELETE status_code=204 elapsed_time=10.307262ms
time=2024-04-13T19:47:23.323+03:00 level=INFO msg="Request status" Path=/banner/1031 Method=DELETE status_code=204 elapsed_time=7.881262ms
time=2024-04-13T19:47:23.324+03:00 level=INFO msg="Request status" Path=/banner/1015 Method=DELETE status_code=204 elapsed_time=8.167352ms
time=2024-04-13T19:47:23.324+03:00 level=INFO msg="Request status" Path=/banner/1001 Method=DELETE status_code=204 elapsed_time=9.409052ms
time=2024-04-13T19:47:23.324+03:00 level=INFO msg="Request status" Path=/banner/1043 Method=DELETE status_code=204 elapsed_time=9.05412ms
time=2024-04-13T19:47:23.324+03:00 level=INFO msg="Request status" Path=/banner/1106 Method=DELETE status_code=204 elapsed_time=6.956933ms
time=2024-04-13T19:47:23.324+03:00 level=INFO msg="Request status" Path=/banner/1048 Method=DELETE status_code=204 elapsed_time=9.394695ms
time=2024-04-13T19:47:23.324+03:00 level=INFO msg="Request status" Path=/banner/1014 Method=DELETE status_code=204 elapsed_time=9.48715ms
time=2024-04-13T19:47:23.325+03:00 level=INFO msg="Request status" Path=/banner/1007 Method=DELETE status_code=204 elapsed_time=10.865778ms
time=2024-04-13T19:47:23.326+03:00 level=INFO msg="Request status" Path=/banner/1076 Method=DELETE status_code=204 elapsed_time=9.505133ms
time=2024-04-13T19:47:23.326+03:00 level=INFO msg="Request status" Path=/banner/1057 Method=DELETE status_code=204 elapsed_time=10.329303ms
time=2024-04-13T19:47:23.327+03:00 level=INFO msg="Request status" Path=/banner/1088 Method=DELETE status_code=204 elapsed_time=7.88595ms
time=2024-04-13T19:47:23.327+03:00 level=INFO msg="Request status" Path=/banner/1041 Method=DELETE status_code=204 elapsed_time=9.7041ms
time=2024-04-13T19:47:23.327+03:00 level=INFO msg="Request status" Path=/banner/1075 Method=DELETE status_code=204 elapsed_time=9.84701ms
time=2024-04-13T19:47:23.327+03:00 level=INFO msg="Request status" Path=/banner/1122 Method=DELETE status_code=204 elapsed_time=8.926699ms
```

Далле представлены логи после 100 операций
```
time=2024-04-13T19:47:23.434+03:00 level=INFO msg="Request status" Path=/banner/2005 Method=DELETE status_code=204 elapsed_time=81.570839ms
time=2024-04-13T19:47:23.435+03:00 level=INFO msg="Request status" Path=/banner/2119 Method=DELETE status_code=204 elapsed_time=33.110925ms
time=2024-04-13T19:47:23.434+03:00 level=INFO msg="Request status" Path=/banner/2639 Method=DELETE status_code=204 elapsed_time=21.978442ms
time=2024-04-13T19:47:23.435+03:00 level=INFO msg="Request status" Path=/banner/1307 Method=DELETE status_code=204 elapsed_time=106.526364ms
time=2024-04-13T19:47:23.435+03:00 level=INFO msg="Request status" Path=/banner/1997 Method=DELETE status_code=204 elapsed_time=85.254929ms
time=2024-04-13T19:47:23.435+03:00 level=INFO msg="Request status" Path=/banner/2177 Method=DELETE status_code=204 elapsed_time=64.646108ms
time=2024-04-13T19:47:23.436+03:00 level=INFO msg="Request status" Path=/banner/2032 Method=DELETE status_code=204 elapsed_time=31.182697ms
time=2024-04-13T19:47:23.435+03:00 level=INFO msg="Request status" Path=/banner/1603 Method=DELETE status_code=204 elapsed_time=101.348318ms
time=2024-04-13T19:47:23.436+03:00 level=INFO msg="Request status" Path=/banner/2824 Method=DELETE status_code=204 elapsed_time=23.114982ms
time=2024-04-13T19:47:23.435+03:00 level=INFO msg="Request status" Path=/banner/2050 Method=DELETE status_code=204 elapsed_time=60.252656ms
time=2024-04-13T19:47:23.435+03:00 level=INFO msg="Request status" Path=/banner/2495 Method=DELETE status_code=204 elapsed_time=37.504838ms
time=2024-04-13T19:47:23.435+03:00 level=INFO msg="Request status" Path=/banner/1938 Method=DELETE status_code=204 elapsed_time=75.47732ms
time=2024-04-13T19:47:23.436+03:00 level=INFO msg="Request status" Path=/banner/2431 Method=DELETE status_code=204 elapsed_time=68.761956ms
time=2024-04-13T19:47:23.435+03:00 level=INFO msg="Request status" Path=/banner/1115 Method=DELETE status_code=204 elapsed_time=109.195224ms
```

Эта проблема описана [README.md](README.md). Из-за нехватки свободных подключений некоторые запросы могут ждать очень долго, некоторым везет, они ждут чуть меньше. Но если листать лог дальше, то видна динамика увеличения времени для каждого запроса вплоть до 350мс.
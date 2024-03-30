# Athens-AR

Ports
www localhost:80
api localhost:8081
www_admin localhost:8082

Api 

figures get

```
const options = {
  method: 'POST',
  headers: {'Content-Type': 'application/json', 'User-Agent': 'insomnia/8.6.1'},
  body: '{"lat":37.97746217615555,"long":23.731130291475555,"radius_m":10}'
};

fetch('http://localhost:8081/figures/get', options)
  .then(response => response.json())
  .then(response => console.log(response))
  .catch(err => console.error(err));
```

figures set

```
const options = {
  method: 'POST',
  headers: {'Content-Type': 'application/json', 'User-Agent': 'insomnia/8.6.1'},
  body: '{"lat":37.97746217616799,"long":23.731130291473484,"marker":"kalas.patt","figure":"kalas.glb"}'
};

fetch('http://localhost:8081/figures/set', options)
  .then(response => response.json())
  .then(response => console.log(response))
  .catch(err => console.error(err));
```
# Projeto  em Go
Visualização da temperatura atual de uma cidade, além de um status de alerta, com base na API [OpenWeatherMap](https://openweathermap.org/current) e baseado [nesse projeto em NodeJs](https://github.com/Ulremberg/API_OpenWeather).

Há três endpoints: root com os endpoints disponíveis; name recebe como parâmetro o nome da cidade; coords recebe como parâmetro latitude e longitude. 

## Regras de negócio:

- Se a umidade do ar estiver abaixo de 30%, deve ser indicado um status de “Umidade baixa”
- Se a temperatura (Celsius) estiver acima de 30 graus, deve ser indicado um status de
"Risco de ensolação”.
- Se a temperatura está entre 10 e 30 graus, e umidade acima de 30%, deve ser indicado um
status de “Nenhum risco eminente”.
- Se a temperatura estiver abaixo de 10 graus, deve ser indicado um status de “Frio Intenso”.
- Status fora das condições anteriores retornam "Status desconhecido".

### Endpoints:

- /name/:name
- /coords/:lat/:lon

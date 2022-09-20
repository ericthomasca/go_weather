# Rusty Weather

A CLI weather app written is Rust.

## Usage

```bash
git clone https://github.com/ericthomasca/go_weather.git
cd go_weather
go get -d ./...
go build
./go_weather {postal code} {country code}
```

## Example

```console
eric@term:~$ ./weather_cli_rust M5T CA
==============================
========  Go Weather  ========
==============================

Weather for Downtown Toronto (Kensington Market / Chinatown / Grange Park) (43.6541, -79.3978)
Last Updated: 2022-09-19 02:44:00

21C (Feels like 22C) Clear
High: 24C  Low: 19C
Wind: 11km/h SSW
Sunrise: 11:00  Sunset: 23:23
```
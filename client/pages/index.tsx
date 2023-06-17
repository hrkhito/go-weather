import { useEffect, useState } from "react";

export default function Home() {
  const [city, setCity] = useState("");
  const [weather, setWeather] = useState(null);
  const [forecast, setForecast] = useState(null);
  const [error, setError] = useState(null);

  const getWeatherData = () => {
    setWeather(null);
    setForecast(null);
    setError(null);

    fetch(`http://localhost:8080/current-weather?city=${city}`)
      .then(response => {
        if (!response.ok) {
          if (response.status === 404) {
            throw Error("そのような都市名は存在しません");
          } else {
            throw Error(response.statusText);
          }
        }
        return response.json();
      })
      .then(data => setWeather(data))
      .catch(error => setError(error.message));

    fetch(`http://localhost:8080/forecast?city=${city}`)
      .then(response => {
        if (!response.ok) {
          if (response.status === 404) {
            throw Error("そのような都市名は存在しません");
          } else {
            throw Error(response.statusText);
          }
        }
        return response.json();
      })
      .then(data => setForecast(data))
      .catch(error => setError(error.message));
  };

  const handleCityChange = (event) => {
    setCity(event.target.value);
  };

  return (
    <div>
      <h1>Weather App</h1>
      <input type="text" value={city} onChange={handleCityChange} />
      <button onClick={getWeatherData}>Get Weather</button>
      {error && <div>Error: {error}</div>}
      {weather && (
        <div>
          <h2>Current Weather</h2>
          <p>Temperature: {weather.temperature}</p>
          <p>Condition: {weather.condition}</p>
        </div>
      )}
      {forecast && (
        <div>
          <h2>Forecast</h2>
          {forecast.map((item, index) => (
            <div key={index}>
              <p>Time: {item.time}</p>
              <p>Temperature: {item.temperature}</p>
              <p>Condition: {item.condition}</p>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

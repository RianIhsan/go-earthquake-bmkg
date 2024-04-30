import React, { useState, useEffect } from 'react';
import axios from 'axios';

const EarthquakeInfo = () => {
  const [earthquakeData, setEarthquakeData] = useState(null);

  useEffect(() => {
    const fetchEarthquakeData = async () => {
      try {
        const response = await axios.get('http://localhost:8080/earthquake');
        setEarthquakeData(response.data);
      } catch (error) {
        console.error('Failed to fetch earthquake data:', error);
      }
    };

    fetchEarthquakeData();
  }, []);

  if (!earthquakeData) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h2>Earthquake Information</h2>
      <p>Kekuatan: {earthquakeData.Magnitude}</p>
      <p>Tanggal: {earthquakeData.DateTime}</p>
      <p>Lokasi: {earthquakeData.Wilayah}</p>
      <p>Kedalaman: {earthquakeData.Kedalaman}</p>
      <p>Dirasakan: {earthquakeData.Dirasakan}</p>
      <p>Potensi: {earthquakeData.Potensi}</p>
      <p>koordinat: {earthquakeData.Coordinates}</p>
      <img src={earthquakeData.Shakemap} alt="Shake Map" />
    </div>
  );
};

export default EarthquakeInfo;

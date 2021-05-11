import React, {useState} from 'react';
import { CountryDropdown, RegionDropdown, CountryRegionData } from 'react-country-region-selector';


function App() {
  const [country, setCountry] = useState('')
  const [region, setRegion] = useState('')
   

  return (
    <div> 
      <CountryDropdown onChange={ val => setCountry(val)} value={country} />
      <RegionDropdown onChange={ val => setRegion(val)} value={region} country={country} />
    </div>  
   
  );
}

export default App;

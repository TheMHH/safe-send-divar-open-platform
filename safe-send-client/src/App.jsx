// src/App.jsx

import React from "react";
import SenderAddressForm from "./components/SenderAddressForm";
import ReceiverAddressForm from "./components/ReceiverAddressForm";

const App = () => {
  return (
    <div>
      <SenderAddressForm />
      <ReceiverAddressForm />
    </div>
  );
};

export default App;

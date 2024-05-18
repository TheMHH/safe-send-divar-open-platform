import { useState } from "react";

const ReceiverAddressForm = () => {
  const [address, setAddress] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await fetch(
        // TODO Set backend endpoint
        "https://your-backend-endpoint.com/api/address",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ receiverAddress: address }),
        },
      );

      if (response.ok) {
        // Handle success - perhaps clear the form, show a success message, etc.
        console.log("Address submitted successfully");
        setAddress("");
      } else {
        // Handle server errors
        console.error("Submission failed");
      }
    } catch (error) {
      // Handle network errors
      console.error("An error occurred", error);
    }
  };

  return (
    <div className="form-container">
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="address">Receiver Address:</label>
          <textarea
            id="address"
            value={address}
            onChange={(e) => setAddress(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="submit-button contrast">
          Submit
        </button>
      </form>
    </div>
  );
};

export default ReceiverAddressForm;

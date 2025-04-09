import { useNavigate, useParams } from "react-router-dom";

const Confirmation = () => {
  const API_URL = import.meta.env.VITE_API_URL;

  const { token = "" } = useParams();
  const navigate = useNavigate();

  const handleConfirm = async () => {
    try {
      const response = await fetch(`${API_URL}/users/activate/${token}`, {
        method: "PUT",
      });

      if (response.ok) {
        // redirect to "/" page
        navigate("/");
      } else {
        // handle error
        alert("Failed to confirm token");
      }
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div>
      <h1>Confirmation Page</h1>

      <button onClick={handleConfirm}>Click to confirm</button>
    </div>
  );
};

export default Confirmation;

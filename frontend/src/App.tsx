import { useState } from "react";
import "./App.css";
import { Screenshot } from "../wailsjs/go/main/App";

function App() {
  const [img, setImg] = useState("");
  const [isProcessing, setIsProcessing] = useState(false);
  const [resultText, setResultText] = useState("Please enter site URL below 👇");
  const [url, setUrl] = useState("");

  const updateUrl = (e: any) => setUrl(e.target.value);

  const getScreenshot = async () => {
    try {
      setIsProcessing(true);
      const base64 = await Screenshot(url);
      setImg(`data:image/png;base64,${base64}`);
      setResultText("Screenshot URL: " + url);
    } catch (err) {
      console.error(err);
    } finally {
      setIsProcessing(false);
    }
  };

  return (
    <div id="App">
      <p>Screenshot</p>
      {isProcessing ? <p>Processing...</p> : (
        <>
          {img && (
            <div>
              <img src={img} width="400" />
            </div>
          )}
          <div id="result" className="result">{resultText}</div>
          <div id="input" className="input-box">
            <input
              id="url"
              className="input"
              onChange={updateUrl}
              autoComplete="off"
              name="input"
              type="text"
            />
            <button className="btn" onClick={getScreenshot}>
              Get screenshot
            </button>
          </div>
        </>
      )}
    </div>
  );
}

export default App;

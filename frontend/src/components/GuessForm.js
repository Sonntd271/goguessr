import React, { useState } from 'react';

const GuessForm = ({ onGuess, onLogout, message, guessNumber }) => {
  const [guess, setGuess] = useState('');

  const handleGuess = () => {
    onGuess(guess);
  };

  return (
    <div>
      <h1 className="h3 mb-3 fw-normal">Guess the number from 1-10!</h1>
      <div className="form-signin m-auto">
        <div className="form-floating my-1">
          <input
            type="number"
            className="form-control"
            id="floating-guess"
            value={guessNumber}
            onChange={(e) => setGuess(e.target.value)}
            placeholder="1"
          />
          <label htmlFor="floating-guess">Guess Number</label>
        </div>
        <span className="text-primary">{message}</span>
        <button className="btn btn-primary w-100 py-2 my-2" onClick={handleGuess}>
          Guess
        </button>
        <button className="btn btn-danger w-100 py-2 my-2" onClick={onLogout}>
          Logout
        </button>
      </div>
    </div>
  );
};

export default GuessForm;

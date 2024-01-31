import React from 'react';
import { useAuth } from './context/AuthContext';
import LoginForm from './components/LoginForm';
import GuessForm from './components/GuessForm';

const App = () => {
  const { token, authenticated, login, logout } = useAuth();
  const [message, setMessage] = React.useState('');
  const [guessNumber, setGuessNumber] = React.useState('');

  const handleLogin = (username, password) => {
    setMessage('');
    login(username, password);
  };

  const handleLogout = () => {
    setMessage('');
    logout();
  };

  const handleGuess = async (guess) => {
    try {
		let body = JSON.stringify({
			guess: parseInt(guessNumber, 10),
			token: token,
		});
		// console.log(body);
		const response = await fetch('http://localhost:8080/guess', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: body,
		});

		if (!response.ok) {
			throw new Error(`Guess failed with status: ${response.status}`);
		}

		const data = await response.text();
		// console.log('Guess successful:', data);
		setMessage(data);
		setGuessNumber('');
	} catch (error) {
		console.error('Guess failed:', error.message);
	}
  };

  return (
    <div className="container-fluid min-vh-100 d-flex align-items-center justify-content-center">
      {!authenticated && <LoginForm onLogin={handleLogin} />}
      {authenticated && (
        <GuessForm
          onGuess={handleGuess}
          onLogout={handleLogout}
          message={message}
          guessNumber={guessNumber}
        />
      )}
    </div>
  );
};

export default App;

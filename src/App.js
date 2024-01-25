import React, { useState } from 'react';

const App = () => {
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');
	const [guessNumber, setGuessNumber] = useState('');
	const [token, setToken] = useState('');
	const [authenticated, setAuthenticated] = useState(false);

	const handleLogin = async () => {
		try {
			var body = JSON.stringify({
				username: username,
				password: password,
			});
			// console.log(body);
			const response = await fetch('http://localhost:8080/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body: body,
			});

			if (!response.ok) {
				throw new Error(`Login failed with status: ${response.status}`);
			}

			const data = await response.json();
			setToken(data.token);
			setAuthenticated(true);
			console.log('Login successful. Token:', data.token);
		} catch (error) {
			console.error('Login failed:', error.message);
		}
	};

	const handleGuess = async () => {
		try {
			var body = JSON.stringify({
				guess: parseInt(guessNumber, 10),
				token: token,
			});
			console.log(body);
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
			console.log('Guess successful:', data);
		} catch (error) {
			console.error('Guess failed:', error.message);
		}
	};

	return (
		<div className="container-fluid min-vh-100 d-flex align-items-center justify-content-center">
			{!authenticated && (
				<div>
					<h1 class="h3 mb-3 fw-normal">Sign in</h1>
					<div className="form-signin m-auto">
						<div className="form-floating my-1">
							<input
								type="text"
								className="form-control"
								id="floating-uname"
								value={username}
								onChange={(e) => setUsername(e.target.value)}
								placeholder="username"
							/>
							<label htmlFor="floating-uname">Username</label>
						</div>
						<div className="form-floating my-1">
							<input
								type="password"
								className="form-control"
								id="floating-password"
								value={password}
								onChange={(e) => setPassword(e.target.value)}
								placeholder="password"
							/>
							<label htmlFor="floating-password">Password</label>
						</div>
						<button
							className="btn btn-primary w-100 py-2 my-2"
							onClick={handleLogin}
						>
							Login
						</button>
					</div>
				</div>
			)}
			{authenticated && (
				<div>
					<h1 class="h3 mb-3 fw-normal">Guess the number!</h1>
					<div className="form-signin m-auto">
						<div className="form-floating my-1">
							<input
								type="text"
								className="form-control"
								id="floating-guess"
								value={guessNumber}
								onChange={(e) => setGuessNumber(e.target.value)}
								placeholder="1"
							/>
							<label htmlFor="floating-guess">
								Guess Number
							</label>
						</div>
						<button
							className="btn btn-primary w-100 py-2 my-2"
							onClick={handleGuess}
						>
							Guess
						</button>
					</div>
				</div>
			)}
		</div>
	);
};

export default App;

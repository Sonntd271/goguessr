import React, { useState } from 'react';
import useGuess from '../hooks/UseGuess';

const GuessForm = ({ onLogout }) => {
	const [guess, setGuess] = useState('');
	const { handleGuess, message } = useGuess();
	// console.log(message)

	const onGuess = () => {
		handleGuess(guess);
		setGuess('')
	};

	let parsedMessage = { message: '' };
	try {
		if (typeof message === 'object' && message.message) {
			parsedMessage = message;
		}
	} catch (error) {
		console.error('Failed to parse JSON:', error.message);
	}
	// console.log(parsedMessage)

	return (
		<div>
			<h1 className="h3 mb-3 fw-normal">Guess the number from 1-10!</h1>
			<div className="form-signin m-auto">
				<div className="form-floating my-1">
					<input
						type="number"
						className="form-control"
						id="floating-guess"
						value={guess}
						onChange={(e) => setGuess(e.target.value)}
						placeholder="1"
					/>
					<label htmlFor="floating-guess">Guess Number</label>
				</div>
				<span className="text-primary">{parsedMessage.message}</span>
				<button
					className="btn btn-primary w-100 py-2 my-2"
					onClick={onGuess}
				>
					Guess
				</button>
				<button
					className="btn btn-danger w-100 py-2 my-2"
					onClick={onLogout}
				>
					Logout
				</button>
			</div>
		</div>
	);
};

export default GuessForm;

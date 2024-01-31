import { useState } from 'react';
import axios from 'axios';
import { useAuth } from '../context/AuthContext';

const useGuess = () => {
	const [message, setMessage] = useState('');
    const { token } = useAuth();

	const handleGuess = async (guessNumber) => {
		try {
			const response = await axios.post(
				'http://localhost:8080/guess',
				{
					guess: parseInt(guessNumber, 10),
				},
				{
					headers: {
						'Content-Type': 'application/json',
						'Authorization': `Bearer ${token}`,
					},
				}
			);

			setMessage(response.data);
		} catch (error) {
			console.error('Guess failed:', error.message);
		}
	};

    // console.log(message)
	return { handleGuess, message };
};

export default useGuess;

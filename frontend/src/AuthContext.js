import React, { createContext, useContext, useState } from 'react';

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
	const [authenticated, setAuthenticated] = useState(false);
	const [token, setToken] = useState('');

	const login = async (username, password) => {
		try {
			const body = JSON.stringify({
				username: username,
				password: password,
			});

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
		} catch (error) {
			console.error('Login failed:', error.message);
		}
	};

	const logout = () => {
		setAuthenticated(false);
		setToken('');
	};

	return (
		<AuthContext.Provider value={{ authenticated, token, login, logout }}>
			{children}
		</AuthContext.Provider>
	);
};

export const useAuth = () => {
	return useContext(AuthContext);
};

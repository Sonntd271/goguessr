import React, { createContext, useContext, useState } from 'react';
import axios from 'axios';

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
	const [authenticated, setAuthenticated] = useState(false);
	const [token, setToken] = useState('');

	const login = async (username, password) => {
		try {
			const response = await axios.post('http://localhost:8080/login', {
				username: username,
				password: password,
			});

			const data = response.data;

			setToken(data.token);
			setAuthenticated(true);

			axios.defaults.headers.common['Authorization'] = `Bearer ${data.token}`;
		} catch (error) {
			console.error('Login failed:', error.message);
		}
	};

	const logout = () => {
		setAuthenticated(false);
		setToken('');

		delete axios.defaults.headers.common['Authorization'];
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

import React from 'react';
import { useAuth } from './context/AuthContext';
import LoginForm from './components/LoginForm';
import GuessForm from './components/GuessForm';

const App = () => {
	const { authenticated, login, logout } = useAuth();

	const handleLogin = (username, password) => {
		login(username, password);
	};

	const handleLogout = () => {
		logout();
	};

	return (
		<div className="container-fluid min-vh-100 d-flex align-items-center justify-content-center">
			{!authenticated ? (
				<LoginForm onLogin={handleLogin} />
			) : (
				<GuessForm
					onLogout={handleLogout}
				/>
			)}
		</div>
	);
};

export default App;

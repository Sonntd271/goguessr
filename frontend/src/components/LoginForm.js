import React, { useState } from 'react';
import { useAuth } from '../context/AuthContext';

const LoginForm = () => {
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');
	const { login } = useAuth();

	const handleLogin = () => {
		login(username, password);
	};

	return (
		<div>
			<h1 className="h3 mb-3 fw-normal">Sign in</h1>
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
	);
};

export default LoginForm;

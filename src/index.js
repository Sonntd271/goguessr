import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import { AuthProvider } from './AuthContext';

ReactDOM.render(
	<React.StrictMode>
		<AuthProvider>
			<App />
		</AuthProvider>
	</React.StrictMode>,
	document.getElementById('root')
);

// TODO 1: Migrate ReactDOM.render
// TODO 2: Display Guess Feedback

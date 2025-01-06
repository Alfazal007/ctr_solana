import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { SignUp } from './components/Signup';
import { SignIn } from './components/Signin';
import Landing from './components/landing';
import UserProvider from './context/UserContext';
import Logout from './components/logout';

export interface User {
	accessToken: string;
	username: string;
	id: string;
	userType: "creator" | "labeller"
}

export default function App() {
	const router = createBrowserRouter([
		{
			path: "/signup",
			element: <SignUp />,
		},
		{
			path: "/signin",
			element: <SignIn />,
		},
		{
			path: "/",
			element: <Landing />
		},
		{
			path: "/logout",
			element: <Logout />
		}
	]);

	return (
		<>
			<UserProvider>
				<RouterProvider router={router} />
			</UserProvider>
		</>
	);

}

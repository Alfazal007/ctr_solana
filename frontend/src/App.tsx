import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { SignUp } from './components/Signup';
import { SignIn } from './components/Signin';
import Landing from './components/landing';
import UserProvider from './context/UserContext';
import Logout from './components/logout';
import CreateNewProject from './components/createNewProject';
import AddImageToProject from './components/addImageToProject';
import TaskCardList from './components/getMyProjects';

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
		},
		{
			path: "/create-project",
			element: <CreateNewProject />
		},
		{
			path: "/add-image/:projectId",
			element: <AddImageToProject />
		},
		{
			path: "/my-projects",
			element: <TaskCardList />
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

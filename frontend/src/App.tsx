import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { SignUp } from './components/Signup';
import { SignIn } from './components/Signin';
import Landing from './components/landing';
import UserProvider from './context/UserContext';
import Logout from './components/logout';
import CreateNewProject from './components/createNewProject';
import AddImageToProject from './components/addImageToProject';
import TaskCardList from './components/getMyProjects';
import EndProject from './components/endProject';
import CreatorSideProject from './components/creatorSideProject';
import VoteProject from './components/voteProject';
import TaskCardListToVote from './components/projectsToVote';
import CreatorTransfer from './components/creatorTransferPage';

import { ConnectionProvider, WalletProvider } from '@solana/wallet-adapter-react';
import {
	WalletModalProvider,
} from '@solana/wallet-adapter-react-ui';
import '@solana/wallet-adapter-react-ui/styles.css';
import WalletConnectPublicKey from './components/publicKeyRegister';

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
		},
		{
			path: "/end-project/:projectId",
			element: <EndProject />
		},
		{
			path: "/project/:projectId",
			element: <CreatorSideProject />
		},
		{
			path: "/solve-task/:projectId",
			element: <VoteProject />
		},
		{
			path: "/projects-to-vote",
			element: <TaskCardListToVote />
		},
		{
			path: "/transfer",
			element: <>
				<ConnectionProvider endpoint={import.meta.env.VITE_SOLANA_ENDPOINT}>
					<WalletProvider wallets={[]} autoConnect>
						<WalletModalProvider>
							<CreatorTransfer />
						</WalletModalProvider>
					</WalletProvider>
				</ConnectionProvider>
			</>
		},
		{
			path: "/add-public-key",
			element: <>
				<ConnectionProvider endpoint={import.meta.env.VITE_SOLANA_ENDPOINT}>
					<WalletProvider wallets={[]} autoConnect>
						<WalletModalProvider>
							<WalletConnectPublicKey />
						</WalletModalProvider>
					</WalletProvider>
				</ConnectionProvider>
			</>
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

import { UserContext } from "@/context/UserContext"
import { WalletDisconnectButton, WalletMultiButton } from "@solana/wallet-adapter-react-ui"
import { useContext, useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import Navbar from "./Navbar"
import axios from "axios"
import { toast } from "@/hooks/use-toast"
import { useConnection, useWallet } from "@solana/wallet-adapter-react"
import { LAMPORTS_PER_SOL, PublicKey, SystemProgram, Transaction } from "@solana/web3.js";

const CreatorTransfer = () => {
	const { user } = useContext(UserContext)
	const navigate = useNavigate()
	const [balance, setBalance] = useState(0);
	const [transferAmount, setTransferAmount] = useState('');
	const wallet = useWallet()
	const { connection } = useConnection()

	useEffect(() => {
		if (!user) {
			navigate("/")
		}
		fetchCreatorBalance()
	}, [])

	async function transfer() {
		if (!wallet.publicKey) {
			return
		}
		let to = "4UH3DAq7tC8SX2GwuJ7P4muZo6DKjmyqUe3oVD4Es1rG";
		const transaction = new Transaction();
		transaction.add(SystemProgram.transfer({
			fromPubkey: wallet.publicKey,
			toPubkey: new PublicKey(to),
			lamports: parseInt(transferAmount) * LAMPORTS_PER_SOL,
		}));
		try {
			await wallet.sendTransaction(transaction, connection);
			toast({ title: "Transfer done" })
		} catch (err) {
			toast({ title: "Issue transferring the amount", variant: "destructive" })
		} finally {
			setTransferAmount("")
		}
	}

	async function fetchCreatorBalance() {
		let balanceResponse;
		try {
			balanceResponse = await axios.get("http://localhost:8000/api/v1/user/balance", { withCredentials: true })
			console.log({ balanceResponse })
			if (balanceResponse.status != 200) {
				toast({
					title: balanceResponse.data
				})
				return
			}
			setBalance(balanceResponse.data)
		} catch (err) {
			toast({
				title: balanceResponse?.data || "Issue fetching the balance"
			})
		}
	}


	return (
		<>
			{user &&
				<>
					<div className="min-h-screen bg-gray-900 text-white">
						<Navbar userType={user.userType} />
						<div className="container mx-auto px-4 py-8">
							<div className="bg-gray-800 rounded-lg shadow-lg p-6 mb-8">
								<h2 className="text-3xl font-bold mb-4">Wallet (Make sure to send from same address as in (Add Public Key page), if not update it and then send</h2>
								<div className="flex space-x-4 mb-6">
									<WalletMultiButton className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded" />
									<WalletDisconnectButton className="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded" />
								</div>
								<div className="mb-8">
									<h3 className="text-2xl font-semibold mb-2">Balance</h3>
									<p className="text-4xl font-bold">{balance} SOL</p>
								</div>
								{
									user.userType == "creator" &&
									<div className="flex items-center space-x-4">
										<input
											type="number"
											value={transferAmount}
											onChange={(e) => setTransferAmount(e.target.value)}
											placeholder="Amount to transfer(in sol)"
											className="flex-grow bg-gray-700 text-white px-4 py-2 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
										/>
										<button
											disabled={!wallet.connected}
											onClick={transfer}
											className={`
    font-bold py-2 px-6 rounded-lg transition duration-300
    ${wallet.connected
													? 'bg-green-500 hover:bg-green-600 text-white'
													: 'bg-gray-400 text-gray-600 cursor-not-allowed'
												}
  `}
										>
											Transfer
										</button>

									</div>
								}
								{
									user.userType != "creator" &&
									<button
										disabled={!wallet.connected}
										onClick={() => { }}
										className={`
    font-bold py-2 px-6 rounded-lg transition duration-300
    ${wallet.connected
												? 'bg-green-500 hover:bg-green-600 text-white'
												: 'bg-gray-400 text-gray-600 cursor-not-allowed'
											}
  `}
									>
										Withdraw
									</button>
								}
							</div>
						</div>
					</div>
				</>
			}
		</>
	);
}

export default CreatorTransfer

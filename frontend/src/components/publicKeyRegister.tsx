import { useContext, useEffect, useState } from 'react'
import { WalletMultiButton, WalletDisconnectButton } from '@solana/wallet-adapter-react-ui'
import { useWallet } from '@solana/wallet-adapter-react'
import { Button } from '@/components/ui/button'
import { UserContext } from '@/context/UserContext'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'
import { toast } from '@/hooks/use-toast'
import bs58 from "bs58"
import { Buffer } from 'buffer'
import Navbar from './Navbar'

export default function WalletConnectPublicKey() {
	const [isRegistered, setIsRegistered] = useState(true)
	const [publicKeyFromDB, setPublicKeyFromDB] = useState("")
	const { publicKey } = useWallet()
	const { user } = useContext(UserContext)
	const navigate = useNavigate()
	const { signMessage } = useWallet()

	const handleStorePublicKey = async () => {
		if (publicKey && user && signMessage) {
			console.log('Storing public key:', publicKey.toBase58())
			const publicKeyToRegister = base58ToBase64(publicKey.toBase58())
			const message = user.id
			const encodedMessage = new TextEncoder().encode(message);
			const signature = await signMessage(encodedMessage);
			const messageToRegister = uint8ArrayToBase64(signature)
			try {
				const addPK = await axios.post("http://localhost:8000/api/v1/user/verify",
					{
						signature: messageToRegister,
						publicKey: publicKeyToRegister
					},
					{ withCredentials: true })
				if (addPK.status == 200) {
					toast({
						title: "Added public key successfully"
					})
					await fetchIsPublicKeyData()
				}
			} catch (err) {
				toast({ title: "Issue registering public key" })
				navigate("/")
			}
		}
	}

	function uint8ArrayToBase64(uint8Array: Uint8Array) {
		let binary = '';
		for (let i = 0; i < uint8Array.length; i++) {
			binary += String.fromCharCode(uint8Array[i]);
		}
		return btoa(binary);
	}

	function base58ToBase64(base58String: string) {
		const byteArray = bs58.decode(base58String);
		const base64String = btoa(String.fromCharCode(...byteArray));
		return base64String;
	}

	function base64ToBase58(base64String: string) {
		const buffer = Buffer.from(base64String, 'base64');
		return bs58.encode(buffer);
	}

	async function fetchIsPublicKeyData() {
		try {
			const publicKeyResponse = await axios.get(`http://localhost:8000/api/v1/user/publicKey`, {
				withCredentials: true
			})
			console.log({ publicKeyResponse })
			if (publicKeyResponse.status != 200) {
				toast({
					title: "Issue fetching the public key, try again later",
					variant: "destructive"
				})
				navigate("/")
				return
			}
			toast({
				title: "Fetched data successfully",
			})
			setIsRegistered(publicKeyResponse.data.isRegistered)
			if (publicKeyResponse.data.publicKey) {
				setPublicKeyFromDB(base64ToBase58(publicKeyResponse.data.publicKey))
			}
		} catch (err) {
			toast({
				title: "Issue fetching the public key, try again later",
				variant: "destructive"
			})
			navigate("/")
		}
	}

	useEffect(() => {
		if (!user) {
			navigate("/signin")
		}
		fetchIsPublicKeyData()
	}, [])

	return (
		<>
			{
				user && <>
					<Navbar userType={user.userType} />
					<div className="min-h-screen bg-gray-900 flex items-center justify-center">
						<div className="bg-gray-800 p-8 rounded-lg shadow-xl max-w-md w-full">
							<h1 className="text-2xl font-bold text-white mb-6 text-center">Wallet Connection</h1>

							<div className="space-y-4">
								<WalletMultiButton className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded w-full" />
								<WalletDisconnectButton className="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded w-full" />
							</div>

							{
								!publicKeyFromDB &&
								<Button
									onClick={handleStorePublicKey}
									className="mt-4 w-full bg-purple-500 hover:bg-purple-600 text-white"
									disabled={!publicKey}
								>
									Store Public Key
								</Button>
							}
							{
								publicKeyFromDB &&
								<h1 className="text-xl font-bold text-white mb-6 text-center">{publicKeyFromDB}</h1>
							}
						</div>
					</div>
				</>
			}
		</>
	)
}


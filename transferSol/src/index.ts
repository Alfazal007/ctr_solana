import {
	Connection,
	PublicKey,
	Keypair,
	Transaction,
	SystemProgram,
	sendAndConfirmTransaction,
	clusterApiUrl,
} from "@solana/web3.js"
import { configDotenv } from "dotenv"
import bs58 from "bs58"
import cors from "cors"
import express, { Request, Response } from "express"

configDotenv()

const connection = new Connection(clusterApiUrl("devnet"))
const privateKeyString = process.env.PRIVATEKEY as string;
const privateKeyArray: Uint8Array = bs58.decode(privateKeyString)
const keypair = Keypair.fromSecretKey(privateKeyArray)

function base64ToBase58(base64String: string): string {
	try {
		const buffer = Buffer.from(base64String, 'base64');
		const base58String = bs58.encode(buffer);
		return base58String;
	} catch (error) {
		return ""
	}
}

async function transfer(amount: number, to: string): Promise<boolean> {
	try {
		const tx = new Transaction().add(
			SystemProgram.transfer({
				lamports: amount,
				fromPubkey: keypair.publicKey,
				toPubkey: new PublicKey(to)
			})
		)
		await sendAndConfirmTransaction(connection, tx, [keypair], {
			commitment: "confirmed"
		})
		console.log("success")
		return true
	} catch (err) {
		console.log("Issue transferring the sol")
		return false
	}
}

const app = express()
app.use(express.json())
app.use(cors())

app.post("/transfer", async (req: Request, res: Response) => {
	if (!req.body.lamports || !req.body.to || !req.body.secret) {
		res.status(400).json({
			message: "Invalid data"
		})
		return
	}

	if (req.body.secret !== process.env.SECRET) {
		res.status(400).json({
			message: "Invalid secret"
		})
		return
	}

	console.log(req.body)

	const isOk = await transfer(req.body.lamports, base64ToBase58(req.body.to))
	if (isOk) {
		res.status(200).json({ message: "done" })
	} else {
		res.status(400).json({ message: "not done" })
	}
})

app.listen(8002, () => {
	console.log("App listening on port 8002")
})

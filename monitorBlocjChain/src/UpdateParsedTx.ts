import web3, { ConfirmedSignatureInfo, ParsedInstruction } from '@solana/web3.js';
import { configDotenv } from 'dotenv';
import axios from 'axios';
import bs58 from "bs58"

configDotenv()

const connection = new web3.Connection(process.env.CONNECTIONURL as string, 'confirmed');
const accountPublicKey = new web3.PublicKey('4UH3DAq7tC8SX2GwuJ7P4muZo6DKjmyqUe3oVD4Es1rG');

export const extractAndUpdateData = async (signature: ConfirmedSignatureInfo[]) => {
	const errorFreeTransactions = signature.filter((sign) => sign.confirmationStatus == "finalized" && sign.err == null).map((sign) => sign.signature);
	const transactions = await connection.getParsedTransactions(errorFreeTransactions);
	for (let i = 0; i < transactions.length; i++) {
		const transaction = transactions[i];
		if (!transaction || !transaction.transaction.message.instructions[2] || transaction.meta?.err) {
			continue;
		}
		const parsedData = transaction.transaction.message.instructions[2] as ParsedInstruction;
		if (parsedData.parsed.type == "transfer") {
			if (parsedData.parsed.info.destination == accountPublicKey) {
				try {
					const receivedAmount = parsedData.parsed.info.lamports;
					const buffer: Uint8Array = bs58.decode(parsedData.parsed.info.source as string);
					// Encode to Base64
					const base64PK = Buffer.from(buffer).toString("base64");
					await axios.post("http://localhost:8000/api/v1/user/update-balance", {
						secret: process.env.SECRET as string,
						address: base64PK,
						lamports: receivedAmount + ""
					})
				} catch (err) {
					console.log(err);
				}
			}
		}
	}
}

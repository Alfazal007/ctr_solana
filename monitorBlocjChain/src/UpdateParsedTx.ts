import web3, { ConfirmedSignatureInfo, ParsedInstruction } from '@solana/web3.js';
import { configDotenv } from 'dotenv';
import { prisma } from './prisma';

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
					console.log({ sentFrom: parsedData.parsed.info.source })
					console.log({ receivedAmount })
					// update the database for prices via api
				} catch (err) {
					console.log(err);
				}
			}
		}
	}
}

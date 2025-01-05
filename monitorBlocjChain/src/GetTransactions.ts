import web3 from '@solana/web3.js';
import { prisma } from './prisma';

const connection = new web3.Connection("https://solana-devnet.g.alchemy.com/v2/216966e_5zv1_orNIihK4tmKzwdFsTJm", 'confirmed');
const address = new web3.PublicKey('4UH3DAq7tC8SX2GwuJ7P4muZo6DKjmyqUe3oVD4Es1rG');

export async function checkActivity(): Promise<web3.ConfirmedSignatureInfo[]> {
	try {
		const lastUsedBlock = await prisma.lastusedblock.findFirst()
		if (!lastUsedBlock) {
			return []
		}
		const signatures = await connection.getSignaturesForAddress(address, { until: lastUsedBlock.lastusedaddress });
		if (signatures.length === 0) {
			return [];
		}
		await prisma.lastusedblock.updateMany({
			data: {
				lastusedaddress: signatures[0].signature
			}
		})
		return signatures
	} catch (error) {
		console.error('Error fetching activity:', error);
		return []
	}
}


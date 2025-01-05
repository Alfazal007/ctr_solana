import { checkActivity } from "./GetTransactions";
import { extractAndUpdateData } from "./UpdateParsedTx";

async function main() {
	//	await prisma.lastusedblock.deleteMany()
	//	await prisma.lastusedblock.create({
	//		data: {
	//			lastusedaddress: ""
	//		}
	//	})
	let signatures = await checkActivity()
	await extractAndUpdateData(signatures)
}

main()

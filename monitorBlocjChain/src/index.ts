import { checkActivity } from "./GetTransactions";
import { extractAndUpdateData } from "./UpdateParsedTx";

async function main() {
	setInterval(async () => {
		let signatures = await checkActivity()
		await extractAndUpdateData(signatures)
	}, 120000)
}

main()

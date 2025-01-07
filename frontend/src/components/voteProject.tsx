import { UserContext } from "@/context/UserContext"
import { toast } from "@/hooks/use-toast"
import axios from "axios"
import { useContext, useEffect, useState } from "react"
import { useNavigate, useParams } from "react-router-dom"

interface ProjectImage {
	secureUrl: string
}

const VoteProject = () => {
	const { user } = useContext(UserContext)
	const navigate = useNavigate()
	const { projectId } = useParams();
	const [imageUrls, setImageUrls] = useState<ProjectImage[]>([])

	async function fetchData() {
		try {
			const projectResponse = await axios.get(`http://localhost:8000/api/v1/project/labeller/${projectId}`, {
				withCredentials: true
			})
			if (projectResponse.status != 200) {
				toast({
					title: "Issue finding the data",
					variant: "destructive"
				})
				navigate("/")
				return
			} else {
				setImageUrls(projectResponse.data)
			}
		} catch (err) {
			toast({
				title: "Issue finding the data",
				variant: "destructive"
			})
			navigate("/")
		}
	}

	useEffect(() => {
		if (!user) {
			navigate("/signin")
			return
		}
		fetchData()
	}, [])

	return (
		<div>VoteProject {JSON.stringify(imageUrls)} </div>
	)
}

export default VoteProject

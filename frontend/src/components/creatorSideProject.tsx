import { toast } from "@/hooks/use-toast"
import axios from "axios"
import { useContext, useEffect, useState } from "react"
import { useNavigate, useParams } from "react-router-dom"
import { Project } from "./addImageToProject"
import { UserContext } from "@/context/UserContext"
import Navbar from "./Navbar"

const CreatorSideProject = () => {
	const [project, setProject] = useState<Project | null>(null)
	const { projectId } = useParams();
	const navigate = useNavigate()
	const { user } = useContext(UserContext)

	async function fetchProject() {
		const response = await axios.get(`http://localhost:8000/api/v1/project/${projectId}`, { withCredentials: true })
		if (response.status != 200) {
			toast({ title: "Issue  fetching project" })
			navigate("/")
			return
		}
		setProject(response.data)
		if (project?.completed == false) {
			toast({ title: "The project is not complete yet" })
			navigate("/")
		}
	}

	useEffect(() => {
		if (!user) {
			navigate("/signin")
			return
		}
		fetchProject()
	}, [])

	return (
		<>
			{
				user && <>
					<Navbar userType={user.userType} />
					<div className="min-h-screen bg-gradient-to-b from-gray-900 to-gray-800 py-12 px-4 sm:px-6 lg:px-8">
						<div className="max-w-7xl mx-auto">
							<div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
								{/*images.map((image, index) => (
						<NumberedImage
							key={index}
							src={image.src}
							alt={image.alt}
							number={index + 1}
						/>
					))
					*/}
							</div>
						</div>
					</div>
				</>
			}
		</>
	)
}

export default CreatorSideProject

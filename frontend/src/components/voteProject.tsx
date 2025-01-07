import { UserContext } from "@/context/UserContext"
import { toast } from "@/hooks/use-toast"
import axios from "axios"
import { useContext, useEffect, useState } from "react"
import { useNavigate, useParams } from "react-router-dom"
import Navbar from "./Navbar"

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

	async function vote() {
		// TODO:: complete this
	}

	return (
		<>
			{
				user && <>
					<Navbar userType={user.userType} />
					<div className="min-h-screen bg-gradient-to-b from-gray-900 to-gray-800 py-12 px-4 sm:px-6 lg:px-8">
						<div className="max-w-7xl mx-auto">
							<div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
								{imageUrls.map((image, index) => (
									<div onClick={vote}>
										<Image
											key={index}
											src={image.secureUrl}
											alt={"Image"}
										/>
									</div>
								))}
							</div>
						</div>
					</div>
				</>
			}
		</>
	)
}

export default VoteProject

interface ImageProps {
	src: string;
	alt: string;
}

function Image({ src, alt }: ImageProps) {
	return (
		<div className="relative">
			<img
				src={src}
				alt={alt}
				className="w-full h-auto rounded-lg object-cover"
			/>
		</div>
	);
}


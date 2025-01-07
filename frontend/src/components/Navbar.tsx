import { useState } from 'react'
import { Menu, X, Plus, Home, DollarSign, LogOut, Folders } from 'lucide-react'
import { Link } from 'react-router-dom'

const Navbar = ({ userType }: { userType: "creator" | "labeller" }) => {
	const [isOpen, setIsOpen] = useState(false)

	const navItems = [
		{ name: userType == "creator" ? 'Create New Project' : 'Work on project', href: '/create-project', icon: Plus },
		{ name: 'Home', href: '/', icon: Home },
		{ name: userType == "creator" ? 'Transfer Money' : 'Take out money', href: '/transfer', icon: DollarSign },
		{ name: 'Logout', href: '/logout', icon: LogOut },
		{ name: userType == "creator" ? "My Projects" : "Projects to vote", href: userType === "creator" ? "/my-projects" : "/projects-to-vote", icon: Folders }
	]

	return (
		<nav className="bg-gray-900">
			<div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
				<div className="flex items-center justify-between h-16">
					<div className="flex items-center">
						<div className="flex-shrink-0">
							<Link to="/" className="text-white font-bold text-xl">
								Home
							</Link>
						</div>
						<div className="hidden md:block">
							<div className="ml-10 flex items-center space-x-4">
								{navItems.map((item) => (
									<Link
										key={item.name}
										to={item.href}
										className="text-gray-300 hover:bg-gray-800 hover:text-white px-3 py-2 rounded-md text-sm font-medium flex items-center"
									>
										<item.icon className="h-5 w-5 mr-2" />
										{item.name}
									</Link>
								))}
							</div>
						</div>
					</div>
					<div className="-mr-2 flex md:hidden">
						<button
							onClick={() => setIsOpen(!isOpen)}
							type="button"
							className="bg-gray-800 inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white"
							aria-controls="mobile-menu"
							aria-expanded="false"
						>
							<span className="sr-only">Open main menu</span>
							{isOpen ? (
								<X className="block h-6 w-6" aria-hidden="true" />
							) : (
								<Menu className="block h-6 w-6" aria-hidden="true" />
							)}
						</button>
					</div>
				</div>
			</div>

			{isOpen && (
				<div className="md:hidden" id="mobile-menu">
					<div className="px-2 pt-2 pb-3 space-y-1 sm:px-3">
						{navItems.map((item) => (
							<Link
								key={item.name}
								to={item.href}
								className="text-gray-300 hover:bg-gray-800 hover:text-white block px-3 py-2 rounded-md text-base font-medium flex items-center"
							>
								<item.icon className="h-5 w-5 mr-2" />
								{item.name}
							</Link>
						))}
					</div>
				</div>
			)}
		</nav>
	)
}

export default Navbar


import "./App.css";
//import AdminLayout from "./containers/AdminLayout";
//import UserLayout from "./containers/UserLayout";
//import Login from "./views/Login";
//import Chat from "./containers/Chat/Chat";
import AdminLayout from "./containers/Layout/AdminLayout";
import Router from "./components/routes/Router";

function App() {
	return (
		<AdminLayout>
			<Router />
		</AdminLayout>
	);
}

export default App;

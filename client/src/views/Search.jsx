import React from "react";
import SearchForm from "../components/forms/SearchForm";

export default function Search() {
	return (
		<>
			<div
				style={{
					padding: 24,
					background: "#fbfbfb",
					border: "1px solid #d9d9d9",
					borderRadius: 2,
				}}
			>
				<SearchForm />
			</div>
			<div className="search-result-list">Search Result List</div>
		</>
	);
}

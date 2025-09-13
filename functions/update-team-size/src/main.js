import { Client, Users } from "node-appwrite";

export default async ({ req, res, log, error }) => {
	const client = new Client()
		.setEndpoint(process.env.APPWRITE_FUNCTION_API_ENDPOINT)
		.setProject(process.env.APPWRITE_FUNCTION_PROJECT_ID)
		.setKey(req.headers["x-appwrite-key"] ?? "");

	try {
	} catch (err) {
		error("Could not list users: " + err.message);
	}

	return res.json({
		message: "Successfully updated team size",
	});
};

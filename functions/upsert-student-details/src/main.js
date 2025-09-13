import { Client } from "node-appwrite";

export default async ({ req, res, log, error }) => {
	const client = new Client()
		.setEndpoint(process.env.APPWRITE_FUNCTION_API_ENDPOINT)
		.setProject(process.env.APPWRITE_FUNCTION_PROJECT_ID)
		.setKey(req.headers["x-appwrite-key"] ?? "");

	try {
	} catch (err) {}

	return res.json({
		message: "Successfully added new student",
	});
};

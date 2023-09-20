import { GetPodMetadata } from "../../wailsjs/go/main/App";

export const getPodMetadata = async (podName, podNamespace) => {
    return await GetPodMetadata(podName, podNamespace);
};

/**
 * returns array of strings
 * @param {string} logs 
 * @returns {Array<string>}
 */
export const splitLogsByNewline = (logs) => {
    return logs.split("\n");
};
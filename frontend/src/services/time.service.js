/**
 * Return human readable time of relative time since the pod has started
 * @param {string} startup 
 * @returns string
 */
export const calculateUptime = (startup) => {
    const now = new Date();
    const since = new Date(startup);

    let result = (now.getTime() - since.getTime()) / 100;
    let suffix = "seconds";

    if (result > 60) {
        result = result / 60;
        suffix = "minutes";
    }

    if (result > 60) {
        result = result / 60;
        suffix = "hours";
    }

    if (result > 24) {
        result = result / 24;
        suffix = "days";
    }

    return `${Math.round(result)} ${suffix}`;
};

export const minutesToSeconds = (minutes) => {
    return minutes * 60;
};
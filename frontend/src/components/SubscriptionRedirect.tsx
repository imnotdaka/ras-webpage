import { useSubscription } from "@/context/SubscriptionContext";
import LoadingPage from "@/pages/LoadingPage";
import { Navigate } from "react-router-dom";

const SubscriptionRedirect = ({ children }: { children: React.ReactNode }) => {
    const { subscription, isSubLoading } = useSubscription();

    if (isSubLoading) {
        return <LoadingPage />;
    }

    if (subscription?.status === "authorized") {
        return <Navigate to="/profile" replace />;
    }

    return <>{children}</>;
};

export default SubscriptionRedirect;
// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.networking.v1alpha1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * ServiceCIDRSpec define the CIDRs the user wants to use for allocating ClusterIPs for Services.
 * 
 */
public final class ServiceCIDRSpecPatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceCIDRSpecPatchArgs Empty = new ServiceCIDRSpecPatchArgs();

    /**
     * CIDRs defines the IP blocks in CIDR notation (e.g. &#34;192.168.0.0/24&#34; or &#34;2001:db8::/64&#34;) from which to assign service cluster IPs. Max of two CIDRs is allowed, one of each IP family. This field is immutable.
     * 
     */
    @Import(name="cidrs")
    private @Nullable Output<List<String>> cidrs;

    /**
     * @return CIDRs defines the IP blocks in CIDR notation (e.g. &#34;192.168.0.0/24&#34; or &#34;2001:db8::/64&#34;) from which to assign service cluster IPs. Max of two CIDRs is allowed, one of each IP family. This field is immutable.
     * 
     */
    public Optional<Output<List<String>>> cidrs() {
        return Optional.ofNullable(this.cidrs);
    }

    private ServiceCIDRSpecPatchArgs() {}

    private ServiceCIDRSpecPatchArgs(ServiceCIDRSpecPatchArgs $) {
        this.cidrs = $.cidrs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceCIDRSpecPatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceCIDRSpecPatchArgs $;

        public Builder() {
            $ = new ServiceCIDRSpecPatchArgs();
        }

        public Builder(ServiceCIDRSpecPatchArgs defaults) {
            $ = new ServiceCIDRSpecPatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cidrs CIDRs defines the IP blocks in CIDR notation (e.g. &#34;192.168.0.0/24&#34; or &#34;2001:db8::/64&#34;) from which to assign service cluster IPs. Max of two CIDRs is allowed, one of each IP family. This field is immutable.
         * 
         * @return builder
         * 
         */
        public Builder cidrs(@Nullable Output<List<String>> cidrs) {
            $.cidrs = cidrs;
            return this;
        }

        /**
         * @param cidrs CIDRs defines the IP blocks in CIDR notation (e.g. &#34;192.168.0.0/24&#34; or &#34;2001:db8::/64&#34;) from which to assign service cluster IPs. Max of two CIDRs is allowed, one of each IP family. This field is immutable.
         * 
         * @return builder
         * 
         */
        public Builder cidrs(List<String> cidrs) {
            return cidrs(Output.of(cidrs));
        }

        /**
         * @param cidrs CIDRs defines the IP blocks in CIDR notation (e.g. &#34;192.168.0.0/24&#34; or &#34;2001:db8::/64&#34;) from which to assign service cluster IPs. Max of two CIDRs is allowed, one of each IP family. This field is immutable.
         * 
         * @return builder
         * 
         */
        public Builder cidrs(String... cidrs) {
            return cidrs(List.of(cidrs));
        }

        public ServiceCIDRSpecPatchArgs build() {
            return $;
        }
    }

}

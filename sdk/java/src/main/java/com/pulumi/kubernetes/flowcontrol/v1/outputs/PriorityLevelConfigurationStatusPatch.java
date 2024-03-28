// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.flowcontrol.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.flowcontrol.v1.outputs.PriorityLevelConfigurationConditionPatch;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class PriorityLevelConfigurationStatusPatch {
    /**
     * @return `conditions` is the current state of &#34;request-priority&#34;.
     * 
     */
    private @Nullable List<PriorityLevelConfigurationConditionPatch> conditions;

    private PriorityLevelConfigurationStatusPatch() {}
    /**
     * @return `conditions` is the current state of &#34;request-priority&#34;.
     * 
     */
    public List<PriorityLevelConfigurationConditionPatch> conditions() {
        return this.conditions == null ? List.of() : this.conditions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PriorityLevelConfigurationStatusPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<PriorityLevelConfigurationConditionPatch> conditions;
        public Builder() {}
        public Builder(PriorityLevelConfigurationStatusPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.conditions = defaults.conditions;
        }

        @CustomType.Setter
        public Builder conditions(@Nullable List<PriorityLevelConfigurationConditionPatch> conditions) {
            this.conditions = conditions;
            return this;
        }
        public Builder conditions(PriorityLevelConfigurationConditionPatch... conditions) {
            return conditions(List.of(conditions));
        }
        public PriorityLevelConfigurationStatusPatch build() {
            final var o = new PriorityLevelConfigurationStatusPatch();
            o.conditions = conditions;
            return o;
        }
    }
}

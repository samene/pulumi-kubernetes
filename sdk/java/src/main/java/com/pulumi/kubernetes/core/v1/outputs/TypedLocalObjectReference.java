// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TypedLocalObjectReference {
    /**
     * @return APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
     * 
     */
    private @Nullable String apiGroup;
    /**
     * @return Kind is the type of resource being referenced
     * 
     */
    private String kind;
    /**
     * @return Name is the name of resource being referenced
     * 
     */
    private String name;

    private TypedLocalObjectReference() {}
    /**
     * @return APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
     * 
     */
    public Optional<String> apiGroup() {
        return Optional.ofNullable(this.apiGroup);
    }
    /**
     * @return Kind is the type of resource being referenced
     * 
     */
    public String kind() {
        return this.kind;
    }
    /**
     * @return Name is the name of resource being referenced
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TypedLocalObjectReference defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String apiGroup;
        private String kind;
        private String name;
        public Builder() {}
        public Builder(TypedLocalObjectReference defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiGroup = defaults.apiGroup;
    	      this.kind = defaults.kind;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder apiGroup(@Nullable String apiGroup) {
            this.apiGroup = apiGroup;
            return this;
        }
        @CustomType.Setter
        public Builder kind(String kind) {
            this.kind = Objects.requireNonNull(kind);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public TypedLocalObjectReference build() {
            final var o = new TypedLocalObjectReference();
            o.apiGroup = apiGroup;
            o.kind = kind;
            o.name = name;
            return o;
        }
    }
}
